# docxtopdf
# golang docx to pdf convertor and S3 package

- Packages dependent on aws sdk go: https://github.com/aws/aws-sdk-go
- For PDF generation, since libreoffice/soffice is sequential,we use collector, dispatcher and worker model to run things in sequence

```
git clone https://github.com/karuppaiah/docxtopdf.git
cd docxtopdf
dep ensure // make sure https://github.com/aws/aws-sdk-go is installed
apt-get install -y libreoffice // If linux env else install libreoffice for mac and put the executables in env PATH
go run main.go // will show a sample of S3 put, pull and del


```
#TODO

- [ ] Unit testing for S3 repository
- [ ] Unit testing for collector, dispatcher and worker pdf generator
