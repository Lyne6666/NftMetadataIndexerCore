# NftMetadataIndexerCore

## Description

This repository houses a framework for generating generative NFT collections with verifiable rarity using on-chain Merkle tree proofs, minimizing gas costs by pre-computing metadata hashes and enabling decentralized minting through a custom ERC-721A implementation.

## Features

- Indexes NFT metadata from multiple EVM-compatible chains using RPC endpoints.
- Implements a configurable caching layer with Redis for optimized query performance.
- Normalizes NFT metadata schemas using a custom data transformation pipeline.
- Supports filtering and sorting of NFT metadata based on user-defined criteria via a GraphQL API.
- Provides a robust event listener that monitors blockchain events for new NFT mints and transfers.
- Integrates with IPFS and Arweave to resolve decentralized storage URIs.
- Automatically retries failed metadata retrieval attempts with exponential backoff.
- Exposes Prometheus metrics for monitoring indexer health and performance.
## Installation

```bash
pip install git+https://github.com/Lyne6666/NftMetadataIndexerCore.git
```

## Usage

```bash
python -m nftmetadataindexercore --verbose
```

## Contributing

We welcome contributions! Here's how to get started:

1. Fork this repository
2. Create a new branch for your feature (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -am 'Add some awesome feature'`)
4. Push to the branch (`git push origin feature/your-feature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.
