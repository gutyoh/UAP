from threading import Lock, current_thread

from hstest.dynamic.input.input_handler import InputHandler
from hstest.dynamic.input.input_mock import Condition
from hstest.dynamic.output.output_handler import OutputHandler
from hstest.dynamic.security.exit_handler import ExitHandler
from hstest.dynamic.security.thread_handler import ThreadHandler
from hstest.exception.outcomes import ErrorWithFeedback
from hstest.testing.execution.program_executor import ProgramExecutor


class SystemHandler:
    __lock = Lock()
    __locked: bool = False
    __locker_thread = None

    @staticmethod
    def set_up():
        SystemHandler._lock_system_for_testing()

        OutputHandler.replace_stdout()
        InputHandler.replace_input()
        ExitHandler.replace_exit()
        ThreadHandler.install_thread_group()

    @staticmethod
    def tear_down():
        SystemHandler._unlock_system_for_testing()

        OutputHandler.revert_stdout()
        InputHandler.revert_input()
        ExitHandler.revert_exit()
        ThreadHandler.uninstall_thread_group()

    @staticmethod
    def _lock_system_for_testing():
        with SystemHandler.__lock:
            if SystemHandler.__locked:
                raise ErrorWithFeedback(
                    "Cannot start the testing process more than once")
            SystemHandler.__locked = True
            SystemHandler.__locker_thread = current_thread()

    @staticmethod
    def _unlock_system_for_testing():
        if current_thread() != SystemHandler.__locker_thread:
            raise ErrorWithFeedback(
                "Cannot tear down the testing process from the other thread")

        with SystemHandler.__lock:
            if not SystemHandler.__locked:
                raise ErrorWithFeedback(
                    "Cannot tear down the testing process more than once")
            SystemHandler.__locked = False
            SystemHandler.__locker_thread = None

    @staticmethod
    def install_handler(program: ProgramExecutor, condition: Condition):
        InputHandler.install_input_handler(program, condition)
        OutputHandler.install_output_handler(program, condition)

    @staticmethod
    def uninstall_handler(program: ProgramExecutor):
        InputHandler.uninstall_input_handler(program)
        OutputHandler.uninstall_output_handler(program)