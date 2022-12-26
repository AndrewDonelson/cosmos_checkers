const BlockchainHost = "http://localhost:1317";

// getNodeInfo returns the node info from the Cosmos Checkers blockchain
// Usage:
//      result = await getNodeInfo();
async function getNodeInfo() {
    try {
        const response = await fetch(BlockchainHost+'/cosmos/base/tendermint/v1beta1/node_info');
        const data = await response.json();
        return data;
    } catch (error) {
        console.error(error);
    }
}

// getLatestBlock returns the latest block from the Cosmos Checkers blockchain
// Usage:
//      result = await getLatestBlock();
async function getLatestBlock() {
    try {
        const response = await fetch(BlockchainHost+'/cosmos/base/tendermint/v1beta1/validatorsets/latest');
        const data = await response.json();
        return data;
    } catch (error) {
        console.error(error);
    }
}
