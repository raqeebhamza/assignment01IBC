package Assignment01IBC

import (
  "fmt"
  "bytes"
  "crypto/sha256"
)

type block struct {
  Hash []byte
  transaction string
  prevPointer *block

}
func InsertBlock(transaction string, chainHead *block) *block {
  b := &block{
    transaction:transaction,
    prevPointer:chainHead,
  }
  b.Hash=b.calculate_Hash()
  return b
}
func (b *block) calculate_Hash()[]byte{
  var data []byte
  if b.prevPointer==nil{
    data = bytes.Join([][]byte{[]byte(b.transaction)},[]byte{})
  }else{
    data = bytes.Join([][]byte{[]byte(b.transaction),b.prevPointer.Hash},[]byte{})
  }
  hash:=sha256.Sum256(data)
  return hash[:]
}
func ListBlocks(chainHead *block){
  for chainHead != nil {
       fmt.Println(chainHead.transaction," :",chainHead.Hash)
       chainHead=chainHead.prevPointer
   }
}
func ChangeBlock(oldstring string,newstring string , chainHead* block){
  for chainHead != nil {
       if chainHead.transaction==oldstring{
         chainHead.transaction=newstring
       }
       chainHead=chainHead.prevPointer
   }
}
func VerifyChain(chainHead *block){
  for chainHead != nil {
       if string(chainHead.calculate_Hash()) != string(chainHead.Hash){
         fmt.Println("Chain contains mellicious blocks, chain has been changed :(")
         return
       }
       chainHead=chainHead.prevPointer
   }
   fmt.Println("Chain is secured and all blocks are verified as trusty :)")
}
