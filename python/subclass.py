
class SuperClass:
    @classmethod
    def run(cls):
        cls._run()

    @classmethod
    def _run(cls):
        raise Exception('must define `_run`')

class SubClass(SuperClass):
    @classmethod
    def _run(cls):
        print('SubClass._run')


SubClass.run()
SuperClass.run()

