use std::env;
use std::fs;

fn main() {
    let args: Vec<String> = env::args().collect();



    println!("In file {}", args[1]);

    let contents = fs::read_to_string(&args[1])
        .expect("Something went wrong reading the file");

    println!("With text:\n{}", contents);
}