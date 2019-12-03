use std::collections::HashMap;
use std::fs::File;
use std::io::BufReader;
use std::io::prelude::*;

fn main() {
    let mut wires = input();

    let a = wires.pop().unwrap();
    let b = wires.pop().unwrap();

    println!("Answer to Part 1 {}", find_closest(&a, &b));
    println!("Answer to Part 2 {}", find_one_with_fewest_steps(&a, &b));
}

type Wire = Vec<Path>;

#[derive(Debug)]
struct Path {
    from: Point,
    to: Point,
    direction: DIRECTION,
    reversed: bool,
    steps: isize
}

#[derive(Debug, Copy, Clone)]
enum DIRECTION {
    VERTICAL,
    HORIZONTAL
}

type Point = (isize, isize);

fn find_one_with_fewest_steps(a: &Wire, b: &Wire) -> isize {
    find_all_intersections(a, b).iter()
        .fold(None, |min, ((_, _), steps)| {
            match min {
                Some(value) if *steps < value => Some(*steps),
                None => Some(*steps),
                _ => min
            }
        }).unwrap()
}

fn find_closest(a: &Wire, b: &Wire) -> isize {
    find_all_intersections(a, b).iter()
        .fold(None, |min, ((x, y), _)| {
            let sum = x.abs() + y.abs();

            match min {
                Some(value) if sum < value => Some(sum),
                None => Some(sum),
                _ => min
            }
        }).unwrap()
}

fn find_all_intersections(a: &Wire, b: &Wire) -> Vec<(Point, isize)> {
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
                      out: &mut Vec<(Point, isize)>) {
    match paths {
        Some(ps) => {
            for p in ps {
                match intersection(path, p) {
                    Some(value) => {
                        out.push((value, calculate_steps(&value, path, p)))
                    },
                    None => { }
                }
            }
        },
        None => { }
    };
}

fn calculate_steps(intersection: &Point, a: &Path, b: &Path) -> isize {
    a.steps + b.steps + correct_steps(intersection, a) + correct_steps(intersection, b)
}

fn correct_steps(point: &Point, path: &Path) -> isize {
    match (path.direction, path.reversed) {
        (DIRECTION::HORIZONTAL, false) => -(path.to.0 - point.0),
        (DIRECTION::HORIZONTAL, true) => path.from.0 - point.0,
        (DIRECTION::VERTICAL, false) => -(path.to.1 - point.1),
        (DIRECTION::VERTICAL, true) => path.from.1 - point.1
    }
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
        .fold((((0, 0), 0), Vec::new()), |((prev, steps), mut paths), x| {
            let (path, point, steps) = as_path(x, prev, steps);

            paths.push(path);

            ((point, steps), paths)
        });

    paths
}

fn as_path(input: &str, from: Point, steps: isize) -> (Path, Point, isize) {
    let (dir, val) = input.split_at(1);
    let value = val.parse::<isize>().unwrap();
    let steps = steps + value;

    match dir {
        "U" =>  {
            let to = (from.0, from.1 + value);

            (Path {
                from: from,
                to: to,
                direction: DIRECTION::VERTICAL,
                reversed: false,
                steps: steps
            }, to, steps)
        },
        "R" => {
            let to = (from.0 + value, from.1);

            (Path {
                from: from,
                to: to,
                direction: DIRECTION::HORIZONTAL,
                reversed: false,
                steps: steps
            }, to, steps)
        },
        "D" => {
            let to = (from.0, from.1 - value);

            (Path {
                from: to,
                to: from,
                direction: DIRECTION::VERTICAL,
                reversed: true,
                steps: steps
            }, to, steps)
        },
        "L" => {
            let to = (from.0 - value, from.1);

            (Path {
                from: to,
                to: from,
                direction: DIRECTION::HORIZONTAL,
                reversed: true,
                steps: steps
            }, to, steps)
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

        let a = as_wire(vec!["R75","D30","R83","U83","L12","D49","R71","U7","L72"]);
        let b = as_wire(vec!["U62","R66","U55","R34","D71","R55","D58","R83"]);

        assert_eq!(find_closest(&a, &b), 159);

        let a = as_wire(vec!["R98","U47","R26","D63","R33","U87","L62","D20","R33","U53","R51"]);
        let b = as_wire(vec!["U98","R91","D20","R16","D67","R40","U7","R15","U6","R7"]);

        assert_eq!(find_closest(&a, &b), 135);
    }

    #[test]
    fn test_find_one_with_fewest_steps() {
        let a = as_wire(vec!["R8","U5","L5","D3"]);
        let b = as_wire(vec!["U7","R6","D4","L4"]);

        assert_eq!(find_one_with_fewest_steps(&a, &b), 30);

        let a = as_wire(vec!["R75","D30","R83","U83","L12","D49","R71","U7","L72"]);
        let b = as_wire(vec!["U62","R66","U55","R34","D71","R55","D58","R83"]);

        assert_eq!(find_one_with_fewest_steps(&a, &b), 610);

        let a = as_wire(vec!["R98","U47","R26","D63","R33","U87","L62","D20","R33","U53","R51"]);
        let b = as_wire(vec!["U98","R91","D20","R16","D67","R40","U7","R15","U6","R7"]);

        assert_eq!(find_one_with_fewest_steps(&a, &b), 410);
    }
}
