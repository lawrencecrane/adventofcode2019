use std::io::BufReader;
use std::io::prelude::*;
use std::fs::File;


fn main() {
    input()
}

fn input() {
    let f = File::open("data/input").unwrap();
    let f = BufReader::new(f);

    for line in f.lines() {
        println!("{}", line.unwrap())
    }
}

