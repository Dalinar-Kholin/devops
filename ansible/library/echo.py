#!/usr/bin/python

from __future__ import (absolute_import, division, print_function)
__metaclass__ = type

from ansible.module_utils.basic import AnsibleModule


def run_module():
    module_args = dict(
        message=dict(type='str', required=True),
        uppercase=dict(type='bool', default=False),
    )

    module = AnsibleModule(
        argument_spec=module_args,
        supports_check_mode=True,
    )

    msg = module.params['message']
    if module.params['uppercase']:
        msg = msg.upper()

    result = {
        'changed': False,
        'original_message': module.params['message'],
        'processed_message': msg,
    }

    if module.check_mode:
        module.exit_json(**result)

    module.exit_json(**result)


def main():
    run_module()


if __name__ == '__main__':
    main()