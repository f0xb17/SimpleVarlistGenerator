import re

def convert_variable(variable):
  """
  Takes a variable name and returns the name without any unnecessary characters.
  It looks for a string in the format "(L.L.variable)" and returns "variable".
  If the string is not in this format, a ValueError is raised with a message describing
  the invalid variable.
  """
  match = re.match(r'\(L\.([L$])\.([a-zA-Z0-9_]+)\)', variable)
  if match:
      return match.group(2)
  raise ValueError(f"Invalid variable: {variable}")

def clean_list(variable_list):
  """
  Takes a list of variable names and returns a new list with all duplicates
  removed. The list is cleaned of unnecessary characters by passing each
  variable through the convert_variable function.
  """
  cleaned_list = []
  for variable in variable_list:
    variable = convert_variable(variable)
    if variable not in cleaned_list:
      cleaned_list.append(variable)
  return cleaned_list

def create_lists(variable_list):
  """
  Takes a list of variable names and separates them into two lists: one for 
  variables containing a '$' in their name and one for the others. Each list 
  is cleaned of duplicates and unnecessary characters. Returns a tuple of two 
  cleaned lists: (varlist, stringvarlist).
  """
  varlist = []
  stringvarlist = []
  for variable in variable_list:
      if '$' in variable:
         stringvarlist.append(variable)
      else:
         varlist.append(variable)

  return clean_list(varlist), clean_list(stringvarlist)
   

def main():
  variabe_list_in = ["(L.L.variable)", "(L.$.variable1)", "(L.$.variable2)", "(L.L.variable3)", "(L.L.variable4)", "(L.$.variable3)", "(L.$.variable1)", "(L.$.variable2)", "(L.$.variable4)", "(L.L.variable)", "(L.L.variable1)", "(L.L.variable2)", "(L.L.variable3)", "(L.L.variable)",] 
  create_lists(variabe_list_in)

  varlist = create_lists(variabe_list_in)[0]
  stringvarlist = create_lists(variabe_list_in)[1]

  print(varlist)
  print(stringvarlist)

if __name__ == "__main__":
  main()
