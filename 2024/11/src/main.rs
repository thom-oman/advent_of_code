use std::{fs, u64};

fn main() {
    let file_path = "./inputs.txt";
    let mut stones = fs::read_to_string(file_path).expect("Could not read file").split(" ").map(|s| {
        s.parse().expect("Cannot parse to u64")
    }).collect::<Vec<u64>>();
    let mut iterations = 25;
    if false { iterations = 75; }

    for i in 0..iterations {
        println!("{i} {:?}", stones.len());
        stones = process_stones(stones);
    }

    println!("{:?}", stones.len())
}

fn process_stones(stones: Vec<u64>) -> Vec<u64> {
    let mut new_stones: Vec<u64> = vec![];
    for st in stones.iter() {
        new_stones.append(&mut process_stone(st));
    }
    new_stones
}

fn process_stone(stone: &u64) -> Vec<u64> {
    let mut new_stones: Vec<u64> = vec![];
    let length = stone.to_string().len();
    if length % 2 == 0 {
        let div: u64 = 10_u64.pow((length / 2).try_into().unwrap());
        let right: u64 = stone % div;
        let left: u64 = (stone - right) / div;
        new_stones.push(left);
        new_stones.push(right);
    } else {
        if *stone == 0 {
            new_stones.push(1);
        } else {
            new_stones.push(stone * 2024);
        }
    }

    new_stones
}
