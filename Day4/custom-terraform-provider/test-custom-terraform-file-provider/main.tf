terraform {
  required_providers {
      localfile = {
         source = "tektutor/file"
      }
   }
}

resource "localfile" "myfile" {
   file_name = "./test1.txt"
   file_content = "Test file"
}
