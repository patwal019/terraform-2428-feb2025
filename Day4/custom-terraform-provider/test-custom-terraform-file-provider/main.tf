terraform {
  required_providers {
      file = {
         source = "tektutor/file"
      }
   }
}

resource "file_localfile" "myfile" {
   file_name = "./test1.txt"
   file_content = "Test file"
}
