const input = `Monkey 0:
Starting items: 71, 56, 50, 73
Operation: new = old * 11
Test: divisible by 13
  If true: throw to monkey 1
  If false: throw to monkey 7

Monkey 1:
Starting items: 70, 89, 82
Operation: new = old + 1
Test: divisible by 7
  If true: throw to monkey 3
  If false: throw to monkey 6

Monkey 2:
Starting items: 52, 95
Operation: new = old * old
Test: divisible by 3
  If true: throw to monkey 5
  If false: throw to monkey 4

Monkey 3:
Starting items: 94, 64, 69, 87, 70
Operation: new = old + 2
Test: divisible by 19
  If true: throw to monkey 2
  If false: throw to monkey 6

Monkey 4:
Starting items: 98, 72, 98, 53, 97, 51
Operation: new = old + 6
Test: divisible by 5
  If true: throw to monkey 0
  If false: throw to monkey 5

Monkey 5:
Starting items: 79
Operation: new = old + 7
Test: divisible by 2
  If true: throw to monkey 7
  If false: throw to monkey 0

Monkey 6:
Starting items: 77, 55, 63, 93, 66, 90, 88, 71
Operation: new = old * 7
Test: divisible by 11
  If true: throw to monkey 2
  If false: throw to monkey 4

Monkey 7:
Starting items: 54, 97, 87, 70, 59, 82, 59
Operation: new = old + 8
Test: divisible by 17
  If true: throw to monkey 1
  If false: throw to monkey 3`;

const monkeys = [];

const operation = (op, current, value) => {
  switch (op) {
    case "+":
      return current + Number(value);
    default:
      return current * Number(value);
  }
};

input.split("\n\n").forEach((singleMonkey) => {
  const monkey = {};
  singleMonkey.split("\n").forEach((monkeyAttr) => {
    if (!monkeyAttr.includes("Monkey")) {
      if (monkeyAttr.includes("Starting")) {
        const attr = monkeyAttr.trim().split(" ");
        monkey.starting = attr
          .splice(2, attr.length)
          .map((el) => el.replace(",", ""))
          .map((el) => Number(el));
      } else if (monkeyAttr.includes("Operation")) {
        const attr = monkeyAttr.trim().split(" ");
        monkey.operation = {
          op: attr[attr.length - 2],
          value: attr[attr.length - 1],
        };
      } else if (monkeyAttr.includes("Test")) {
        const attr = monkeyAttr.trim().split(" ");
        monkey.test = { value: attr[attr.length - 1] };
      } else if (monkeyAttr.includes("true")) {
        const attr = monkeyAttr.trim().split(" ");
        monkey.test = { ...monkey.test, true: attr[attr.length - 1] };
      } else {
        const attr = monkeyAttr.trim().split(" ");
        monkey.test = { ...monkey.test, false: attr[attr.length - 1] };
      }
    }
  });
  monkeys.push({ ...monkey, inspections: 0 });
});

for (let i = 0; i < 20; i++) {
  for (let j = 0; j < monkeys.length; j++) {
    monkeys[j].starting.forEach(el => {
      monkeys[j].inspections++;
      let opRes = operation(
        monkeys[j].operation.op,
        el,
        monkeys[j].operation.value === "old"
          ? el
          : monkeys[j].operation.value
      );

      opRes = Math.floor(opRes / 3);
      const testCase = opRes % Number(monkeys[j].test.value);
      if (testCase === 0) {
        monkeys[Number(monkeys[j].test.true)].starting.push(opRes);
      } else {
        monkeys[Number(monkeys[j].test.false)].starting.push(opRes);
      }
    })
    monkeys[j].starting = [];
  }
}

monkeys.sort((a,b) => b.inspections - a.inspections);
console.log("Solution for Day 11 Part 1: ", monkeys[0].inspections * monkeys[1].inspections)

