package loader

import (
	celerywide "github.com/go-celery/celery-wite"
	"github.com/go-celery/celery-wite/enum"
	"github.com/go-celery/celery-wite/errs"
	"github.com/go-celery/celery-wite/factory"
	"github.com/go-celery/celery-wite/funcs"
	"github.com/go-errors/errors"
	"github.com/gocelery/gocelery"
)

type Loader struct {
	transmitFromContexts []funcs.TransmitFromContext
	retarder             celerywide.Retarder
	brokerLoadItem       celerywide.LoadItem
	producerCoreClient   *celerywide.CoreClient
	backend              gocelery.CeleryBackend
	consumerCoreClients  map[enum.QueueName]*celerywide.CoreClient
	logger               celerywide.Logger
}

func New(loadItems []celerywide.LoadItem) (celerywide.Loader, error) {
	return new(Loader).loads(loadItems)
}

func (l *Loader) GetRetarder() celerywide.Retarder {
	return l.retarder
}

func (l *Loader) GetLogger() celerywide.Logger {
	return l.logger
}

func (l *Loader) GetProducerCoreClient() *celerywide.CoreClient {
	return l.producerCoreClient
}

func (l *Loader) GetTransmitFromContexts() []funcs.TransmitFromContext {
	return l.transmitFromContexts
}

func (l *Loader) CreateConsumerCoreClient(queue enum.QueueName, numWorkers int) (*celerywide.CoreClient, error) {
	if consumerCoreClient := l.consumerCoreClients[queue]; consumerCoreClient != nil {
		return consumerCoreClient, nil
	}

	consumerBroker, err := factory.CreateConsumerBroker(l.brokerLoadItem.GetLoadItemConfig(), queue)
	if err != nil {
		return nil, err
	}

	consumerCoreClient, err := l.newCeleryClient(consumerBroker, l.backend, numWorkers)
	if err != nil {
		return nil, err
	}

	if consumerCoreClient == nil {
		return nil, errors.New("broker load item not found")
	}

	l.consumerCoreClients[queue] = consumerCoreClient
	return consumerCoreClient, nil
}

func (l *Loader) GetAllCoreClients() []*celerywide.CoreClient {
	all := make([]*celerywide.CoreClient, 0)
	for _, celeryClient := range l.consumerCoreClients {
		all = append(all, celeryClient)
	}

	all = append(all, l.producerCoreClient)
	return all
}

func (l *Loader) loads(loadItems []celerywide.LoadItem) (*Loader, error) {
	l.consumerCoreClients = make(map[enum.QueueName]*celerywide.CoreClient)
	for _, loadItem := range loadItems {
		if !loadItem.IsLoad() {
			continue
		}

		if err := loadItem.Verify(); err != nil {
			return l, err
		}

		switch loadItem.GetType() {
		case enum.BrokerType:
			if err := l.loadBroker(loadItem); err != nil {
				return l, err
			}
			break
		case enum.BackendType:
			if err := l.loadBackend(loadItem); err != nil {
				return l, err
			}
			break
		case enum.RetarderType:
			if err := l.loadRetarder(loadItem); err != nil {
				return l, err
			}
			break
		case enum.TransmitType:
			if err := l.loadTransmit(loadItem); err != nil {
				return l, err
			}
			break
		case enum.LogType:
			if err := l.loadLogger(loadItem); err != nil {
				return l, err
			}
			break
		default:
			return l, errs.NewError("load item not found: %d", loadItem.GetType())
		}
	}

	return l, nil
}

func (l *Loader) newCeleryClient(broker celerywide.Broker, backend gocelery.CeleryBackend, numWorkers int) (*celerywide.CoreClient, error) {
	celeryClient, err := gocelery.NewCeleryClient(broker, backend, numWorkers)
	if err != nil {
		return nil, err
	}
	queue, err := broker.GetQueueName()
	if err != nil {
		return nil, err
	}

	coreClient := &celerywide.CoreClient{
		CeleryClient: celeryClient,
		Queue:        queue,
	}

	coreClient.StartWorker()
	return coreClient, nil
}

func (l *Loader) loadBroker(loadItem celerywide.LoadItem) error {
	if l.brokerLoadItem != nil {
		return errors.New("loader broker load item can only be one")
	}

	l.brokerLoadItem = loadItem
	producerBroker, err := factory.CreateProducerBroker(l.brokerLoadItem.GetLoadItemConfig())
	if err != nil {
		return err
	}

	producerCoreClient, err := l.newCeleryClient(producerBroker, l.backend, 0)
	if err != nil {
		return err
	}

	l.producerCoreClient = producerCoreClient
	return nil
}

func (l *Loader) loadBackend(loadItem celerywide.LoadItem) (err error) {
	if l.backend, err = factory.CreateBackend(loadItem.GetLoadItemConfig()); err != nil {
		return err
	}

	return nil
}

func (l *Loader) loadRetarder(loadItem celerywide.LoadItem) (err error) {
	if l.retarder, err = factory.CreateRetarder(loadItem.GetLoadItemConfig()); err != nil {
		return err
	}

	return
}

func (l *Loader) loadTransmit(loadItem celerywide.LoadItem) error {
	transmitFromContexts, err := factory.CreateTransmit(loadItem.GetLoadItemConfig())
	if err != nil {
		return err
	}

	l.transmitFromContexts = transmitFromContexts
	return nil
}

func (l *Loader) loadLogger(loadItem celerywide.LoadItem) error {
	logger, err := factory.CreateLogger(loadItem.GetLoadItemConfig())
	if err != nil {
		return err
	}

	l.logger = logger
	return nil
}
