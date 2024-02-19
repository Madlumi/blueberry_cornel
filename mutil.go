package m;
import (
   "fmt"
)
func Er(msg string, err error) bool{
   if err != nil { 
      fmt.Print( fmt.Errorf("ERROR: %s: %v", msg, err));
      return true;
   }
   return false;
}
