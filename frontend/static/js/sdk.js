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

// 
// getCheckersSystemInfo returns the current state from the Cosmos Checkers blockchain
// Usage:
//      result = await getCheckersSystemInfo();
async function getCheckersSystemInfo() {
    try {
        const response = await fetch(BlockchainHost+'/checkers/checkers/system_info');
        const data = await response.json();
        return data;
    } catch (error) {
        console.error(error);
    }
}

// getCheckersStoredGames returns the stored games from the Cosmos Checkers blockchain
// Usage:
//      result = await getCheckersStoredGames();
async function getCheckersStoredGames() {
    try {
        const response = await fetch(BlockchainHost+'/checkers/checkers/stored_game');
        const data = await response.json();
        return data;
    } catch (error) {
        console.error(error);
    }
}

// getCheckersStoredGame returns the stored game of the provided index from the Cosmos Checkers blockchain
// Usage:
//      result = await getCheckersStoredGame(index);
async function getCheckersStoredGame(index) {
    try {
        const response = await fetch(BlockchainHost+'/checkers/checkers/stored_game/'+index);
        const data = await response.json();
        return data;
    } catch (error) {
        console.error(error);
    }
}