use std::fs::File;
use std::io::BufReader;
use std::io::prelude::*;

fn main() {
    println!("{:?}", input())
}

fn input() -> Vec<String> {
    let f = File::open("data/input").unwrap();
    let f = BufReader::new(f);

    let mut out = Vec::new();

    for line in f.lines() {
        let data = line.unwrap();
        out.push(data);
    }

    out
}
