# Source Code Documentation

Since Golang is not supported in Doxygen, we use a tool called Sphinx and several plugins.

To get the necessary tools, use the following commands (Golang and Python3 needed):

```
go get -u github.com/readthedocs/godocjson # make sure you have godocjson in PATH
pip3 install sphinx sphinx-autoapi sphinx-rtd-theme
pip3 install https://github.com/martykan/sphinxcontrib-golangdomain/archive/refs/heads/master.tar.gz
```

To generate the documentation, run

```
make html
```

Then the output will be found in \_build/html
