use std::{f32::consts::E, fs};
use regex::Regex;

use valkyrja::{Exception, throw_value_exception, throw_file_exception, get_exception_message, raise_info, raise_error};

#[derive(Copy, Clone)]
enum VariableType {
    StringVariable,
    Variable,
}

struct Variable {
    name: String,
    variable_type: VariableType,
}

impl Variable {
    fn get_name(&self) -> String {
        return String::from(&self.name).clone();
    }

    fn get_type(&self) -> &VariableType {
        return &self.variable_type;
    }

    fn store(&self, storage: &mut Vec<String>) {
        storage.push(self.get_name());
    }
}

fn return_variable_name(value: &str) -> String {
    let re = Regex::new(r"\([LS]\.[L$]\.([a-zA-Z0-9_]+)\)").unwrap();
    return re.captures(value).unwrap().get(1).unwrap().as_str().to_string();
}

fn collect_variables(file_path: &str) -> Result<Vec<Variable>, Exception> {
    if file_path == "" {
        return Err(throw_file_exception("File Path should not be empty!"));
    }
    
    let mut collected_variables: Vec<Variable> = Vec::new();

    let mut contents = fs::read_to_string(file_path);
    match contents {
        Err(e) => println!("{}", raise_error(&e.to_string())),
        Ok(values) => {
            for value in values.lines() {
                match value.find("$") {
                    Some(_) => collected_variables.push(Variable {
                        name: return_variable_name(&value),
                        variable_type: VariableType::StringVariable
                    }),
                    None => collected_variables.push(Variable {
                        name: return_variable_name(&value),
                        variable_type: VariableType::Variable
                    }),
                }
            }

        }
    }
    if collected_variables.len() == 0 {
        return Err(throw_value_exception("Couldn't find any Variable. Maybe the file is wrong?"));
    }

    Ok(collected_variables)
}

fn main() {
    let result = collect_variables("./src/test.txt");
    match result {
        Ok(vector) => {
            for variable in vector {
                println!("{}", variable.get_name());
            }
        },
        Err(e) => println!("{}", get_exception_message(e))
    }
}
