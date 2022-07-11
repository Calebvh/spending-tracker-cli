use std::env;
use std::fs;
use chrono::{DateTime, NaiveDate, NaiveDateTime, NaiveTime};
use chrono::format::ParseError;
//Verify Commit test attempt2
fn main() {
    let args: Vec<String> = env::args().collect();



    println!("In file {}", args[1]);

    let contents = fs::read_to_string(&args[1])
        .expect("Something went wrong reading the file");
    let rows = contents.split("\n");
    let mut i = 0;
    for line in rows{
        if i > 2{
            return;
        }
        let cols: Vec<&str> = line.split(",").collect();
        i+=1;

        match parse_date(&cols[0].replace("\"", "")) {
            Ok(description) => println!("{}", description),
            Err(err) => println!("Parsing Date: {} {}", &cols[0], err),
        }
    }
}

fn print_type_of<T>(_: &T) {
    println!("{}", std::any::type_name::<T>())
}



fn parse_date(date_str : &str) -> Result<NaiveDate, ParseError>{

    let date_only = NaiveDate::parse_from_str(date_str, "%m/%d/%Y")?;

    Ok(date_only)

}