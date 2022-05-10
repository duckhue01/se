## dependency management
when our code import an external package, its dependency graph is augmented not only with the imported package but also with its set of transitive dependencies - that is, any other packages required by packages that we import. As our project larger in size, it become necessary to efficiently manage the version of all our dependencies to ensure that change in upstream transitive dependencies do not cause unexpected side effects to our own programs.



### what's all the fuss about software versioning?

having a sane versioning system in place makes it possible to do the following:
1. validate that a particular piece of software can be used as a safe drop-in replacement for an older of software that we are using as part of our production systems.
2. pin down each dependency of our applications to a particular package version.

### semantic versioning
semantic versioning is a widely popular system for describing software versions in a way that  makes it quite straightforward for the intended software users to figure out which version are safe to upgrade to and which versions contain breaking change and therefore require development  effort and time when upgrading


- MAJOR.MINOR.PATCH

1. patch: is incremented whenever a backward-compatible bug fix is applied to the code
2. minor: is incremented when new functionality gets added to a package
3. major: is incremented when a breaking change is made
### Applying semantic versioning to go package

### managing the source code for multiple package version
one thing that may have you as odd about the previous section is the fact that while i keep going on about the publishing version x.y.z of the weather package, the section content 
1. single repository with versioned folders
pros:
- the user of a single repository for all version makes maintenance easier, as package authors can work on each version of the package in isolation.
- the repository always contains the latest release for each package version. The end users oft he package can use a single command to get/update all versions of the package
cons:
- code duplication
- end user don't know the newest version
1. single repository - multiple branches

### vendoring - the good, the bad, the ugly

#### is vendoring always a good idea?
#### strategies and tools for vendoring dependencies
1. dep tool
2. go module
3. fork
...


### wrapping up
...
### questions

1. why is software versioning importance?
2. what does a semantic version look like and when are its individual components incremented?
3. which component of a package's semantic version would increment in the following cases?  
- a anew api is introduced?
- an existing api is modified and a new, required parameter is added to it
- a fix for a security bug is committed
4. name some alternative versioning schemes what we could use besides semantic versioning
5. what are the pros and cons of vendoring?
6. name some of the differences between te dep tool and Go modules