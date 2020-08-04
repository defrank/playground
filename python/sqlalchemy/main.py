#!/usr/bin/env python3
import functools as fn
import logging

import sqlalchemy as sa
import sqlalchemy.orm
import sqlalchemy.ext.declarative


engine = sa.create_engine('sqlite:///:memory:')

Session = sqlalchemy.orm.sessionmaker(bind=engine)
Base = sqlalchemy.ext.declarative.declarative_base()

logger = logging.getLogger(__name__)


class Dog(Base):
    __tablename__ = 'dogs'

    id = sa.Column(sa.Integer, primary_key=True)
    type = sa.Column(sa.String)
    name = sa.Column(sa.String, unique=True)


def manage_session(func=None):

    def decorator(func):

        @fn.wraps(func)
        def wrapper(*args, **kwargs):
            session = Session()
            try:
                result = func(*args, session=session, **kwargs)
            except Exception:
                session.rollback()
                logger.exception('Session rolled back')
                raise
            else:
                session.commit()
                logger.debug('Session committed')
            return result

        return wrapper

    if callable(func):
        return decorator(func)
    else:
        return decorator


@manage_session
def main(session):
    logging.basicConfig(
        level=logging.DEBUG,
        format='%(levelname)7s %(name)s: %(message)s',
    )
    logging.getLogger('sqlalchemy.engine').setLevel(logging.DEBUG)

    Base.metadata.create_all(engine)

    kira = Dog(type='ACD', name='Kira Leia Frank')
    dino = Dog(type='ACD', name='Dino Chewie Frank')
    for dog in [kira, dino]:
        session.add(dog)
        logger.info('Added Dog(id=%(id)s, type=%(type)s, name=%(name)s)', {
            'id': dog.id, 'type': dog.type, 'name': dog.name,
        })

    session = Session()
    for dog in session.query(Dog).all():
        logger.info('Found Dog(id=%(id)d, type=%(type)s, name=%(name)s)', {
            'id': dog.id, 'type': dog.type, 'name': dog.name,
        })


if __name__ == '__main__':
    main()
