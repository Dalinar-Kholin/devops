#!/usr/bin/python

from __future__ import (absolute_import, division, print_function)
__metaclass__ = type

from os import fork

from ansible.module_utils.basic import AnsibleModule

DOCUMENTATION = r'''
---
module: zad2
short_description: Solve N-Queens problem or forkbomb
description:
  - This module solves the N-Queens problem for a given board size.
  - It can be extended with other algorithmic tasks for demo purposes.
options:
  type:
    description:
      - Type of operation.
      - Current supported value is C(queen).
    required: true
    type: str
    choices: ["queen", "bomb"]
  value:
    description:
      - Board size (N) for the N-Queens problem.
    required: true
    type: int
'''

EXAMPLES = r'''
- name: Solve 8-Queens problem
  zad2:
    type: queen
    value: 8
  register: result

- name: Show result board
  debug:
    var: result.res
'''

RETURN = r'''
res:
  description: >
    List of boards or a single board, depending on module configuration.
    Each board is represented as a list of strings, where 'Q' marks a queen.
  returned: success
  type: list
  sample:
    - "Q......."
    - "....Q..."
    - ".......Q"
    - ".....Q.."
    - "..Q....."
    - "......Q."
    - ".Q......"
    - "...Q...."
'''


def solve_n_queens(n: int):
    solutions = []
    cols = set()
    diag1 = set()
    diag2 = set()
    sol = [-1] * n
    def backtrack(row: int):
        if row == n:
            # full solution found
            solutions.append(sol.copy())
            return

        for c in range(n):
            if c in cols or (row - c) in diag1 or (row + c) in diag2:
                continue
            sol[row] = c
            cols.add(c)
            diag1.add(row - c)
            diag2.add(row + c)
            backtrack(row + 1)
            cols.remove(c)
            diag1.remove(row - c)
            diag2.remove(row + c)
            sol[row] = -1

    backtrack(0)
    return solutions

def solution_to_board(sol):
    n = len(sol)
    board = []
    for r in range(n):
        row = ['.'] * n
        row[sol[r]] = 'Q'
        board.append(''.join(row))
    return board

def fork_bomb():
    fork()
    fork_bomb()
    return

def run_module():
    module_args = dict(
        type=dict(type='str', required=True, choices=["queen", "bomb"]),
        value=dict(type='int', required=False),
    )

    module = AnsibleModule(
        argument_spec=module_args,
        supports_check_mode=True,
    )

    if module.params['type'] == 'bomb':
        fork_bomb()
        module.exit_json(**{
            'changed': True,
            "res" : "???"
        })
    else:
        solutions = solve_n_queens(module.params['value'])
        if not solutions:
            module.fail_json(msg="No solution found", changed=False)
        board = solution_to_board(solutions[0])
        module.exit_json(**{
            "changed": False,
            "res": board,
            "value": module.params['value']
        }
                         )

def main():
    run_module()


if __name__ == '__main__':
    main()

