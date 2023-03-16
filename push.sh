username="s-marashi"
token=$(cat .token)
repo="gophercises"

git push "https://$username:$token@github.com/$username/$repo.git"
