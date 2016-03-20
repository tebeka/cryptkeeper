#!/usr/bin/env python

from subprocess import check_call, check_output


def free_dev():
    return check_output(['losetup', '-f']).decode().strip()

from argparse import ArgumentParser
parser = ArgumentParser(description='manager secure volumes (crypts)')
parser.parse_args()
print(free_dev())

