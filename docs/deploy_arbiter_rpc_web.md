## deploy arbiter

Note: The server needs to open port 8080 and have deployed the loan-arbiter by the steps of `deploy_loan_arbiter.md`

1. Log in to the server use command:

   ```shell
   ssh ssh -L 8080:127.0.0.1:8080 user@your_server_ip
   ```

2. Enter the home directory
   ```shell
   cd ~
   ```

3. Download deploy script
   ```shell
   wget https://download.bel2.org/loan-arbiter/deploy_arbiter_rpc_web.sh
   ```

4. Script permission changes
   ```shell
   chmod a+x deploy_arbiter_rpc_web.sh
   ```

5. Execute deploy script
   ```shell
   ./deploy_arbiter_rpc_web.sh
   ```

6. Check rpc and web status 

   check deploy script status succeed or failed.

7. Logs

   detailed rpc log: ~/loan_arbiter/data/logs/rpc.log
   detailed web log: ~/loan_arbiter/data/logs/web.log

## kill rpc and web

   ```shell
   pkill -x "arbiter_rpc"
   pkill -x "arbiter_web"
   ```

## restart rpc and web

   ```shell
   cd ~/loan_arbiter
   ./arbiter_rpc --gf.gcfg.file=config.yaml  > ~/loan_arbiter/data/logs/arbiter_rpc.log 2>&1 &
   ./arbiter_web --gf.gcfg.file=config.yaml  > ~/loan_arbiter/data/logs/arbiter_web.log 2>&1 &
   ```
