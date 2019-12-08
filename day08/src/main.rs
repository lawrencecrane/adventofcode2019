use std::fs::File;
use std::io::BufReader;
use std::io::prelude::*;

fn main() {
    println!("{}", input())
}

fn input() -> String {
    let f = File::open("data/input").unwrap();
    let f = BufReader::new(f);

    f.lines().next().unwrap().unwrap()
}

// #[cfg(test)]
// mod tests {
//     use super::*;

//     #[test]
//     fn test_fn() {
//     }
// }
