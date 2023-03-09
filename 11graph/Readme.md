#### 图


```
protoc --go_opt=paths=source_relative \
--go_opt=Mtest_two_proto/test_proto/test.proto=va.zhst.com/internal/pkg/zhstscreen/test_two_proto/test_proto \
 --go_opt=Mtest_two_proto/a.proto=/internal/pkg/zhstscreen/zhstscreen \
 --go_out=plugins=grpc:/code/internal/pkg/zhstscreen test_two_proto/test_proto/test.proto  test_two_proto/a.proto
```