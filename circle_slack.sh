#!/bin/bash

string=$(git log --format=%s -n 1)
if [[ $string =~ .*[fix].* ]] ;
then
  echo "No Slack because commit message have prefix [fix]"
  exit 0
fi

msg="\
\`$CIRCLE_PROJECT_REPONAME\`
Halo, saya $CIRCLE_USERNAME mau meminta review untuk PR berikut @cart-dev
Github   : $CIRCLE_PULL_REQUEST
<!channel>
"

#if [ "$CIRCLE_BRANCH" != "development" ] && [ "$CIRCLE_BRANCH" != "staging" ] ; then
#    curl -X POST --data-urlencode "payload={\"channel\": \"#testrobby2\", \"username\": \"\", \"text\": \"$msg\",\"link_names\": 1, \"icon_emoji\": \":davidjenggot:\"}" https://hooks.slack.com/services/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
#    echo "\nalready slack"
#else
#    echo "no slack"
#fi
