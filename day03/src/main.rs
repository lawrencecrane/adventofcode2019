use std::io::BufReader;
use std::io::prelude::*;
use std::fs::File;

fn main() {
    input()
}

type Wire = Vec<Path>;

#[derive(Debug)]
struct Path {
    from: Point,
    to: Point,
    direction: DIRECTION
}

#[derive(Debug)]
enum DIRECTION {
    VERTICAL,
    HORIZONTAL
}

type Point = (isize, isize);

fn as_wire(input: Vec<&str>) -> Wire {
    let (_, paths) = input.iter()
        .fold(((0, 0), Vec::new()), |(prev, mut paths), x| {
            let path = as_path(x, prev);
            let point = path.to;

            paths.push(path);

            (point, paths)
        });

    paths
}

fn as_path(input: &str, from: Point) -> Path {
    let (dir, val) = input.split_at(1);
    let value = val.parse::<isize>().unwrap();

    match dir {
        "U" => Path {
            from: from,
            to: (from.0, from.1 + value),
            direction: DIRECTION::VERTICAL
        },
        "R" => Path {
            from: from,
            to: (from.0 + value, from.1),
            direction: DIRECTION::HORIZONTAL
        },
        "D" => Path {
            from: (from.0, from.1 - value),
            to: from,
            direction: DIRECTION::VERTICAL
        },
        "L" => Path {
            from: (from.0 - value, from.1),
            to: from,
            direction: DIRECTION::HORIZONTAL
        },
        _ => panic!("Unknown direction")
    }
}

fn input() {
    let f = File::open("data/input").unwrap();
    let f = BufReader::new(f);

    let mut wires = Vec::new();

    for line in f.lines() {
        let data = line.unwrap();
        let input: Vec<&str> = data.split(',').collect();
        wires.push(as_wire(input));
    }

    println!("{:?}", &wires);
}

