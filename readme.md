generate .go from .proto:
```bash
protoc --proto_path=./ --go_out=./ --go_opt=paths=source_relative kv.proto types/types.proto
```
