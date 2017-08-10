Train Server

vagrant ssh pro4ma
sudo su www-data
cd; . setenv
gg
alias goi='go install'
alias runq='mysql -utrain -pkaluma train -e'

clear && goi && trainServer whereami -33.9420123 18.4913511

http://192.168.50.5:8083/whereami?latitude=-33.9420123&longitude=18.4913511
http://target.pro4ma.net:8083/whereami?latitude=-33.9420123&longitude=18.4913511
curl -H "Content-Type: application/json" --data '{"username":"xyz","password":"xyz"}' "http://localhost:8083/whereami?latitude=-33.9420123&longitude=18.4913511"

