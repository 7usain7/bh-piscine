curl -s https://learn.reboot01.com/assets/superhero/all.json | jq '.name | select(.id == 170)'
curl -s https://learn.reboot01.com/assets/superhero/all.json | jq '.powerstats.power | select(.id == 170)'
curl -s https://learn.reboot01.com/assets/superhero/all.json | jq '.appearance.gender | select(.id == 170)'

