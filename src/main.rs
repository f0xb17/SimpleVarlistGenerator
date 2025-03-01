use valkyrja::{Exception, throw_value_exception, throw_file_exception, get_exception_message};

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

fn collect_variables(file_path: &str) -> Result<Vec<Variable>, Exception> {
    if file_path == "" {
        return Err(throw_file_exception("File Path should not be empty!"));
    }
    
    let collected_variables: Vec<Variable> = Vec::new();

    if collected_variables.len() == 0 {
        return Err(throw_value_exception("Couldn't find any Variable. Maybe the file is wrong?"));
    }

    Ok(collected_variables)
}

fn main() {
    let result = collect_variables("./");
    match result {
        Ok(vector) => println!("nothing"),
        Err(e) => println!("{}", get_exception_message(e))
    }
}
