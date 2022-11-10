 package arch

 import (
   "runtime"
   "testing"
 )

 func TestDependant(t * testing.T) {
   t.Parallel()
   if runtime.GOARCH == "amd64" {
     t.Skip("Não funciona em arquitetura amd64")
   }

   t.Fail()
 }
