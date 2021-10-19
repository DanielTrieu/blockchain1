cd /home/daniel/mychain
git add -A
git commit -a -m "update node0"
eval "$(ssh-agent -s)"
ssh-add ~/.ssh/strapi_ssh
git push -u git@github.com:DanielTrieu/blockchain1.git main