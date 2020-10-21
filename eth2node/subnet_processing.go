package eth2node

import (
	"context"
	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"go.uber.org/zap"
)

func (n *Eth2Node) subnetValidator(index DASSubnetIndex) pubsub.ValidatorEx {
	return func(ctx context.Context, p peer.ID, msg *pubsub.Message) pubsub.ValidationResult {
		// TODO subnet validation
		return pubsub.ValidationAccept
	}
}

func (n *Eth2Node) handleSubnet(index DASSubnetIndex, sub *pubsub.Subscription) {
	{
		msg, err := sub.Next(n.subProcesses.ctx)
		if err != nil {
			if err == n.subProcesses.ctx.Err() {
				return
			}
			n.log.With(zap.Error(err)).With("subnet", index).Error("failed to read from subnet subscription")
			sub.Cancel()
			return
		}
		n.log.With("from", msg.ReceivedFrom, "length", len(msg.Data)).Debug("received message")
	}
}
