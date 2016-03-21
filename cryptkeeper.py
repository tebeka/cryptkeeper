#!/usr/bin/env python

from subprocess import check_call, check_output
from os import makedirs, environ
from os.path import dirname, expanduser, isdir
import sys

if sys.version_info[:2] < (3, 0):
    from ConfigParser import ConfigParser
else:
    from configparser import ConfigParser


rc_dir = expanduser('~/.config/cryptkeeper')


def load_config():
    cp = ConfigParser()
    cp.read(['%s/crypts.ini' % rc_dir])
    crypts = {}  # name -> path
    for name in cp.sections():
        path = cp.get(name, 'path')
        if path.startswith('~'):
            path = expanduser(path)
        crypts[name] = path

    return crypts


def free_dev():
    return check_output(['losetup', '-f']).decode().strip()


def tc_name(name):
    return '%s/tcs/%s.tc' % (rc_dir, name)


def mapper_name(name):
    return '/dev/mapper/%s.tc' % name


def ensure_dir(fname):
    root = dirname(fname)
    if not isdir(root):
        makedirs(root)


def create(name, path, size):
    dev = free_dev()
    tc = tc_name(dev)
    mapper = mapper_name(name)
    vol = vol_name(name)
    user = environ['USER']

    ensure_dir(tc)
    ensure_dir(vol)

    cmds = [
        ['fallocate', '-l', size, tc],
        ['losetup', dev, tc],
        ['tcplay', '-c', '-d', dev, '-a', 'whirlpool', '-b', 'AES-256-XTS'],
        ['tcplay', '-m', tc, '-d', dev],
        ['mkfs.ext4', mapper],
        ['mount', mapper, vol],
        ['chown', '-R', '%s:%s' % (user, user), vol],
    ]
    for cmd in cmds:
        check_call(cmd)

    update_config(name, path)


if __name__ == '__main__':
    from argparse import ArgumentParser

    def runargs(fn):
        def runner(args):
            return fn(**vars(args))
        return runner

    parser = ArgumentParser(description='manager secure volumes (crypts)')
    subps = parser.add_subparsers(help='sub-command help')
    crp = subps.add_parser('create', help='create crypt')
    crp.add_argument('name', help='crypt name')
    crp.add_argument('path', help='crypt volume path')
    crp.add_argument('--size', help='crypt size (e.g. 2G)', default='1G')
    crp.set_defaults(func=runargs(create))
    args = parser.parse_args()

    args.func(args)
