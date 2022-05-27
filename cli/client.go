package cli

import (
	"context"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	entropyv1beta1 "go.buf.build/odpf/gwv/odpf/proton/odpf/entropy/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const timeout = 2

func createConnection(ctx context.Context, host string) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	}

	return grpc.DialContext(ctx, host, opts...)
}

func createClient(cmd *cobra.Command) (entropyv1beta1.ResourceServiceClient, func(), error) {
	c, err := loadConfig(cmd)
	if err != nil {
		return nil, nil, err
	}

	host := c.Service.Host + ":" + strconv.Itoa(c.Service.Port)

	dialTimeoutCtx, dialCancel := context.WithTimeout(cmd.Context(), time.Second*timeout)
	conn, err := createConnection(dialTimeoutCtx, host)
	if err != nil {
		dialCancel()
		return nil, nil, err
	}

	cancel := func() {
		dialCancel()
		conn.Close()
	}

	client := entropyv1beta1.NewResourceServiceClient(conn)
	return client, cancel, nil
}
