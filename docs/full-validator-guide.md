# Full Validator Step-by-Step Guide

This guide walks through provisioning a brand new server to run a **Nectar** validator using the `validator-assistant` TUI.  Each section expands on the tasks the assistant automates so you can understand and reproduce them manually if needed.

> **Note**: Commands below assume an Ubuntu based host and a bash shell.  Adjust paths and package managers as required for your environment.

## 1. Prepare the Host

1. Update the base system and install build tools:
   ```bash
   sudo apt update && sudo apt upgrade -y
   sudo apt install -y build-essential curl git jq make systemd
   ```
2. Install Go (version 1.20 or newer):
   ```bash
   curl -LO https://go.dev/dl/go1.20.7.linux-amd64.tar.gz
   sudo tar -C /usr/local -xzf go1.20.7.linux-amd64.tar.gz
   echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.profile
   source ~/.profile
   go version  # verify installation
   ```
3. (Optional) Create a dedicated user to run the validator:
   ```bash
   sudo useradd -m -s /bin/bash nectar
   sudo usermod -aG sudo nectar
   su - nectar
   ```

## 2. Build the Nectar Node

1. Fetch the node source and build the `nectard` binary:
   ```bash
   git clone https://github.com/nectarchain/nectar
   cd nectar
   make install   # installs nectard to $HOME/go/bin
   ```
2. Confirm the binary is available:
   ```bash
   which nectard
   nectard version --long
   ```

## 3. Initialise the Chain

1. Create the configuration directory and initialise the node:
   ```bash
   nectard init <MONIKER> --chain-id nectar-1
   ```
2. Download the genesis and address book files:
   ```bash
   curl -L https://snapshots.nectarchain.io/genesis.json > ~/.nectard/config/genesis.json
   curl -L https://snapshots.nectarchain.io/addrbook.json > ~/.nectard/config/addrbook.json
   ```
3. Configure peers, seeds and pruning (full node shown):
   ```bash
   sed -i 's/^seeds *=.*/seeds = "seed1:26656,seed2:26656"/' ~/.nectard/config/config.toml
   sed -i 's/^persistent_peers *=.*/persistent_peers = "peer1:26656,peer2:26656"/' ~/.nectard/config/config.toml
   sed -i 's/^pruning *=.*/pruning = "default"/' ~/.nectard/config/app.toml
   ```

## 4. Configure with Validator Assistant

1. Obtain this repository and build the TUI:
   ```bash
   cd ~
   git clone https://github.com/nectarchain/validator-assistant
   cd validator-assistant
   go build -o validator-assistant
   ```
2. Launch the assistant:
   ```bash
   ./validator-assistant
   ```
3. Follow the on-screen prompts to choose the node type (full node, archive, etc.), set chain values and validator metadata.  The assistant shows a live YAML preview.
4. Save the configuration when prompted.  By default it writes to `~/.nectard/validator.yaml`.

## 5. Start the Service

1. Create a `systemd` service to run the node:
   ```bash
   sudo tee /etc/systemd/system/nectard.service <<'EOT'
   [Unit]
   Description=Nectar Node
   After=network-online.target

   [Service]
   User=%i
   ExecStart=$(which nectard) start
   Restart=on-failure
   LimitNOFILE=65535

   [Install]
   WantedBy=multi-user.target
   EOT
   ```
2. Enable and start the service:
   ```bash
   sudo systemctl daemon-reload
   sudo systemctl enable --now nectard
   journalctl -u nectard -f   # view logs
   ```

## 6. Create the Validator

After the node has fully synced, submit a create-validator transaction using the values from your saved YAML file:
```bash
nectard tx staking create-validator \
  --amount 1000000unectar \
  --pubkey "$(nectard tendermint show-validator)" \
  --moniker "<MONIKER>" \
  --chain-id nectar-1 \
  --commission-rate 0.10 \
  --commission-max-rate 0.20 \
  --commission-max-change-rate 0.01 \
  --min-self-delegation 1 \
  --from <WALLET> \
  --gas auto --gas-adjustment 1.4
```

## 7. Enable Metrics and Grafana

1. Turn on Prometheus metrics in `config.toml`:
   ```bash
   sed -i 's/^prometheus *=.*/prometheus = true/' ~/.nectard/config/config.toml
   ```
2. Install Grafana:
   ```bash
   sudo apt install -y apt-transport-https software-properties-common wget
   wget -q -O - https://packages.grafana.com/gpg.key | sudo apt-key add -
   sudo add-apt-repository "deb https://packages.grafana.com/oss/deb stable main"
   sudo apt update
   sudo apt install -y grafana
   sudo systemctl enable --now grafana-server
   ```
3. Visit `http://<server-ip>:3000` to log in (default user/pass `admin/admin`).
4. Add Prometheus as a data source pointing to the node's metrics port (`http://localhost:26660`) and import a Nectar dashboard.

## 8. Maintenance Tips

- Update the node with `git pull && make install`.
- Use snapshots to speed up sync: `curl -L https://snapshots.nectarchain.io/latest.tar.lz4 | lz4 -dc - | tar -xf - -C ~/.nectard`.
- For archive mode, set `pruning = "nothing"` in `app.toml`.

With these steps complete your server is running a fully configured Nectar validator with monitoring in place.
