#!/usr/bin/env python3
import logging

from bottle import request, route, run


logger = logging.getLogger(__name__)


@route('/headers/<name>')
def echo_header(name):
    header = request.headers.get(name)
    logger.debug('Header %s=%r', name, header)
    return header


if __name__ == '__main__':
    logging.basicConfig(level=logging.DEBUG)
    run(host='localhost', port=8080)
