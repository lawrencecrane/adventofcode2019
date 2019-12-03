use std::collections::HashMap;
use std::fs::File;
use std::io::BufReader;
use std::io::prelude::*;

fn main() {
    let mut wires = input();

    let a = wires.pop().unwrap();
    let b = wires.pop().unwrap();

    println!("Answer to Part 1 {}", find_closest(&a, &b));
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

fn find_closest(a: &Wire, b: &Wire) -> isize {
    find_all_intersections(a, b).iter()
        .fold(None, |min, (x, y)| {
        let sum = x.abs() + y.abs();

        match min {
            Some(value) if sum < value => Some(sum),
            None => Some(sum),
            _ => min
        }
    }).unwrap()
}

fn find_all_intersections(a: &Wire, b: &Wire) -> Vec<Point> {
    let (horizontal, vertical) = split_wire(&b);

    a.iter()
        .fold(Vec::new(), |mut xs, path| {
            match path.direction {
                DIRECTION::HORIZONTAL => {
                    for x in path.from.0..path.to.0 {
                        find_intersections(path, vertical.get(&x), &mut xs)
                    }
                },
                DIRECTION::VERTICAL => {
                    for y in path.from.1..path.to.1 {
                        find_intersections(path, horizontal.get(&y), &mut xs)
                    }
                }
            };

            xs
        })
}

fn find_intersections(path: &Path,
                      paths: Option<&Vec<&Path>>,
                      out: &mut Vec<Point>) {
    match paths {
        Some(ps) => {
            for p in ps {
                match intersection(path, p) {
                    Some(value) => out.push(value),
                    None => { }
                }
            }
        },
        None => { }
    };
}

fn intersection(a: &Path, b: &Path) -> Option<Point> {
    match a.direction {
        DIRECTION::HORIZONTAL => intersection_helper(a, b),
        DIRECTION::VERTICAL => intersection_helper(b, a)
    }
}

fn intersection_helper(horizontal: &Path, vertical: &Path) -> Option<Point> {
    if horizontal.from.0 < vertical.from.0 && vertical.from.0 < horizontal.to.0 &&
        vertical.from.1 < horizontal.from.1 && horizontal.from.1 < vertical.to.1 {
        Some((vertical.from.0, horizontal.from.1))
    } else {
        None
    }
}

fn split_wire(wire: &Wire) -> (HashMap<isize, Vec<&Path>>, HashMap<isize, Vec<&Path>>) {
    wire.iter()
        .fold((HashMap::new(), HashMap::new()), |(mut hor, mut ver), path| {
            match path.direction {
                DIRECTION::HORIZONTAL => {
                    let paths = hor.entry(path.from.1).or_insert(Vec::new());
                    paths.push(path);
                    (hor, ver)
                },
                DIRECTION::VERTICAL => {
                    let paths = ver.entry(path.from.0).or_insert(Vec::new());
                    paths.push(path);
                    (hor, ver)
                }
            }
        })
}

fn as_wire(input: Vec<&str>) -> Wire {
    let (_, paths) = input.iter()
        .fold(((0, 0), Vec::new()), |(prev, mut paths), x| {
            let (path, point) = as_path(x, prev);

            paths.push(path);

            (point, paths)
        });

    paths
}

fn as_path(input: &str, from: Point) -> (Path, Point) {
    let (dir, val) = input.split_at(1);
    let value = val.parse::<isize>().unwrap();

    match dir {
        "U" =>  {
            let to = (from.0, from.1 + value);

            (Path {
                from: from,
                to: to,
                direction: DIRECTION::VERTICAL
            }, to)
        },
        "R" => {
            let to = (from.0 + value, from.1);

            (Path {
                from: from,
                to: to,
                direction: DIRECTION::HORIZONTAL
            }, to)
        },
        "D" => {
            let to = (from.0, from.1 - value);

            (Path {
                from: to,
                to: from,
                direction: DIRECTION::VERTICAL
            }, to)
        },
        "L" => {
            let to = (from.0 - value, from.1);

            (Path {
                from: to,
                to: from,
                direction: DIRECTION::HORIZONTAL
            }, to)
        },
        _ => panic!("Unknown direction")
    }
}

fn input() -> Vec<Wire> {
    let f = File::open("data/input").unwrap();
    let f = BufReader::new(f);

    let mut wires = Vec::new();

    for line in f.lines() {
        let data = line.unwrap();
        let input: Vec<&str> = data.split(',').collect();
        wires.push(as_wire(input));
    }

    wires
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_find_closest() {
        let a = as_wire(vec!["R8","U5","L5","D3"]);
        let b = as_wire(vec!["U7","R6","D4","L4"]);

        assert_eq!(find_closest(&a, &b), 6);
    }
}
