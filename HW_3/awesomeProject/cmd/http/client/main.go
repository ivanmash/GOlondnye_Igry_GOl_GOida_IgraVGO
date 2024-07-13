package main

import (
	dto2 "awesomeProject/accounts/http/dto"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

type Command struct {
	Port    int
	Host    string
	Cmd     string
	Name    string
	Amount  int
	NewName string
}

func (c *Command) Do() error {
	switch c.Cmd {
	case "create":
		return c.create()
	default:
		return fmt.Errorf("unknown command: %s", c.Cmd)
	}
}

func (c *Command) create() error {
	panic("implement me")
}

func main() {
	portVal := flag.Int("port", 1323, "server port")
	hostVal := flag.String("host", "0.0.0.0", "server host")
	cmdVal := flag.String("cmd", "", "command to execute")
	nameVal := flag.String("name", "", "name of account")
	amountVal := flag.Int("amount", 0, "amount of account")
	newnameVal := flag.String("newname", "", "newname of account")

	flag.Parse()

	cmd := Command{
		Port:    *portVal,
		Host:    *hostVal,
		Cmd:     *cmdVal,
		Name:    *nameVal,
		Amount:  *amountVal,
		NewName: *newnameVal,
	}

	if err := do(cmd); err != nil {
		panic(err)
	}
}

func do(cmd Command) error {
	switch cmd.Cmd {
	case "create":
		if err := create(cmd); err != nil {
			return fmt.Errorf("create account failed: %w", err)
		}

		return nil
	case "get":
		if err := get(cmd); err != nil {
			return fmt.Errorf("get account failed: %w", err)
		}

		return nil
	case "delete":
		if err := delete(cmd); err != nil {
			return fmt.Errorf("delete account failed: %w", err)
		}

		return nil
	case "patch":
		if err := patch(cmd); err != nil {
			return fmt.Errorf("patch account failed: %w", err)
		}

		return nil
	case "change":
		if err := change(cmd); err != nil {
			return fmt.Errorf("change account failed: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unknown command %s", cmd.Cmd)
	}
}

func get(cmd Command) error {
	resp, err := http.Get(
		fmt.Sprintf("http://%s:%d/account?name=%s", cmd.Host, cmd.Port, cmd.Name),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read body failed: %w", err)
		}

		return fmt.Errorf("resp error %s", string(body))
	}

	var response dto2.GetAccountResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("json decode failed: %w", err)
	}

	fmt.Printf("response account name: %s and amount: %d", response.Name, response.Amount)

	return nil
}

func create(cmd Command) error {
	request := dto2.CreateAccountRequest{
		Name:   cmd.Name,
		Amount: cmd.Amount,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/create", cmd.Host, cmd.Port),
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusCreated {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func delete(cmd Command) error {
	request := dto2.DeleteAccountRequest{
		Name: cmd.Name,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://%s:%d/account/delete", cmd.Host, cmd.Port), bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	defer func() {
		_ = req.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func patch(cmd Command) error {
	var request dto2.PatchAccountRequest

	request.Name = cmd.Name
	request.Amount = cmd.Amount

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	req, err := http.NewRequest("PATCH",
		fmt.Sprintf("http://%s:%d/account/patch", cmd.Host, cmd.Port),
		bytes.NewReader(data),
	)
	if err != nil {
		return fmt.Errorf("create put request failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("send put request failed: %w", err)
	}

	if err != nil {
		return fmt.Errorf("http put failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func change(cmd Command) error {
	var request dto2.ChangeAccountRequest

	request.Name = cmd.Name
	request.NewName = cmd.NewName

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	req, err := http.NewRequest("PUT",
		fmt.Sprintf("http://%s:%d/account/change", cmd.Host, cmd.Port),
		bytes.NewReader(data),
	)
	if err != nil {
		return fmt.Errorf("create put request failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("send put request failed: %w", err)
	}

	if err != nil {
		return fmt.Errorf("http put failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}
