package main

import (
	pb "awesomeProject/proto"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Command struct {
	Port    int
	Host    string
	Cmd     string
	Name    string
	NewName string
	Amount  int
}

func main() {
	portVal := flag.Int("port", 4567, "server port")
	hostVal := flag.String("host", "localhost", "server host")
	cmdVal := flag.String("cmd", "", "command to execute")
	nameVal := flag.String("name", "", "name of account")
	newNameVal := flag.String("newname", "", "new name")
	amountVal := flag.Int("amount", 0, "amount of account")

	flag.Parse()

	cmd := Command{
		Port:    *portVal,
		Host:    *hostVal,
		Cmd:     *cmdVal,
		Name:    *nameVal,
		NewName: *newNameVal,
		Amount:  *amountVal,
	}

	if err := cmd.Do(); err != nil {
		panic(err)
	}
}

func (c *Command) Do() error {
	conn, err := grpc.NewClient(c.Host+":"+fmt.Sprint(c.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return fmt.Errorf("grpc connection failed: %w", err)
	}

	defer conn.Close()

	switch c.Cmd {
	case "create":
		if err := c.create(conn); err != nil {
			return fmt.Errorf("create account failed: %w", err)
		}
		return nil

	case "get":
		if err := c.get(conn); err != nil {
			return fmt.Errorf("get account failed: %w", err)
		}
		return nil

	case "delete":
		if err := c.delete(conn); err != nil {
			return fmt.Errorf("delete account failed: %w", err)
		}
		return nil

	case "patch":
		if err := c.patch(conn); err != nil {
			return fmt.Errorf("patch account failed: %w", err)
		}
		return nil

	case "change":
		if err := c.change(conn); err != nil {
			return fmt.Errorf("change account name failed: %w", err)
		}
		return nil

	default:
		return fmt.Errorf("unknown command %s", c.Cmd)
	}
}

func (c *Command) create(conn *grpc.ClientConn) error {
	client := pb.NewAccountsClient(conn)

	request := &pb.CreateAccountRequest{Name: c.Name, Amount: int32(c.Amount)}

	_, err := client.CreateAccount(context.Background(), request)
	if err != nil {
		return fmt.Errorf("grpc create account failed: %w", err)
	}

	return nil
}

func (c *Command) delete(conn *grpc.ClientConn) error {
	client := pb.NewAccountsClient(conn)

	request := &pb.DeleteAccountRequest{
		Name: c.Name,
	}

	_, err := client.DeleteAccount(context.Background(), request)

	if err != nil {
		return fmt.Errorf("grpc delete account failed: %w", err)
	}

	return nil
}

func (c *Command) get(conn *grpc.ClientConn) error {
	client := pb.NewAccountsClient(conn)

	request := &pb.GetAccountRequest{Name: c.Name}

	response, err := client.GetAccount(context.Background(), request)

	if err != nil {
		return fmt.Errorf("grpc get account failed: %w", err)
	}

	fmt.Printf("Name: %s, Amount: %d\n", response.Name, response.Amount)

	return nil
}

func (c *Command) patch(conn *grpc.ClientConn) error {
	client := pb.NewAccountsClient(conn)

	request := &pb.PatchAccountRequest{Name: c.Name, Amount: int32(c.Amount)}

	_, err := client.PatchAccount(context.Background(), request)

	if err != nil {
		return fmt.Errorf("grpc patch account failed: %w", err)
	}

	return nil
}

func (c *Command) change(conn *grpc.ClientConn) error {
	client := pb.NewAccountsClient(conn)

	request := &pb.ChangeAccountRequest{
		Name:    c.Name,
		Newname: c.NewName,
	}

	_, err := client.ChangeAccount(context.Background(), request)

	if err != nil {
		return fmt.Errorf("grpc change account name failed: %w", err)
	}

	return nil
}
