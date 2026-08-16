# go-ctx-detach

子调用脱离父取消

internal/rpcx/client.go 的 Do：内部 WithTimeout 基于 context.Background 而非传入的 parent
