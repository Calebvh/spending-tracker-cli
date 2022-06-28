use std::env;
use std::fs;
use crate::chrono::{DateTime};
use crate::chrono::format::ParseError;

fn main() {
    let args: Vec<String> = env::args().collect();



    println!("In file {}", args[1]);

    let contents = fs::read_to_string(&args[1])
        .expect("Something went wrong reading the file");
    let rows = contents.split("\n");
    let mut i = 0;
    for line in rows{
        if i > 5{
            break;
        }
        let cols = contents.split(",");
        i+=1;
        for item in cols{
            println!("{}", item);
        }
    }

    //println!("With text:\n{}", contentsArr[0]);
}