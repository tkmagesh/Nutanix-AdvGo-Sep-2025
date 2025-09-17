1. Define the contract
    - Service contract
    - Operation contracts
    - Data / Message contracts
2. Generate the proxy & stub
    Stub -> interface (contract)
    Proxy -> concrete type (struct)
3. (Service) Implement the service using the stub contract
4. (Service) Host the service
5. (Client) Use the proxy to communicate to the server