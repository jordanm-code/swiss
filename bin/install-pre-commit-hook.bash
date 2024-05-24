#!/usr/bin/env bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
PRE_COMMIT_HOOK_FILE=$SCRIPT_DIR/../.git/hooks/pre-commit

if [ -f $PRE_COMMIT_HOOK_FILE ]; then
  if grep -q gen_readme.bash $PRE_COMMIT_HOOK_FILE; then
    echo "Pre-commit hook already exists."
    exit 0
  fi
else
  echo "#!/usr/bin/env sh" > $PRE_COMMIT_HOOK_FILE
  chmod +x $PRE_COMMIT_HOOK_FILE
fi

`cat $SCRIPT_DIR/pre-commit.template | tail -n +2 >> $PRE_COMMIT_HOOK_FILE`
