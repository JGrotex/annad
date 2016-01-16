# development
This document describes the development environment and its processes.

### setup
At first, clone the repository.
```
git clone git@github.com:xh3b4sd/anna.git
```

Fetching all dependencies works with this.
```
make goget
```

Just running the code without installing the binary can be achieved using the
go tools.
```
make gorun
```

Cleanup the workspace can be done with this.
```
make goclean
```

### pull requests

###### commits
Pull requests are only accepted and merged when there is only one commit to be
merged. This means contributers need to squash their commits before. This can
be done with the following command.
```
git rebase -i master
```

###### diary
Pull requests are only accepted and merged when there is some sort of process
documentation. Goal of this is to keep track of influences and events that
drove development and decisions. All ideas and walkthroughs are precious and
good to know.