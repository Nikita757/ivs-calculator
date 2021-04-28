project = 'IVS Calculator'
copyright = '2021, FIT VUT Team'
author = 'FIT VUT Team'

# -- General configuration

extensions = ['autoapi.extension', 'sphinx_rtd_theme', 'sphinxcontrib.golangdomain']

exclude_patterns = ['_build', 'Thumbs.db', '.DS_Store']

# -- AutoAPI

autoapi_type = 'go'

autoapi_dirs = ['../pkg/mathfunc']

def prepare_docstring(string):
  print(string)
  lines = string.split("\n")
  output = ""
  for line in lines:
    line = line.strip()
    if line.startswith("*"):
      line = line[1:]
    line = line.strip()

    # Process Javadoc
    if line.startswith("@"):
      first_space = line.index(" ")
      second_space = line.index(" ", first_space+1)
      line = ":" + line[1:second_space] + ":" + line[second_space:]
    output += line + "\n\n"
  return output

def autoapi_prepare_jinja_env(jinja_env):
  jinja_env.filters["format_docstring"] = prepare_docstring

autoapi_template_dir = './_templates'

# -- Options for HTML output

html_theme = 'sphinx_rtd_theme'