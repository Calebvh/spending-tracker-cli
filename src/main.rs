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
        if i > 5{
            return;
        }
        let cols = line.split(",");
        i+=1;
        for item in cols{
            println!("{}", item);
            parseDate(item);
        }
    }

    //println!("With text:\n{}", contentsArr[0]);
}


fn parseDate(dateStr : &str) -> Result<NaiveDate, ParseError>{

    let date_only = NaiveDate::parse_from_str(dateStr, "%m/%d/%Y")?;
    println!("{}", date_only);


    Ok(date_only)

}