import os
from enum import Enum
from Structures import *

PathType = Enum("File", "Directory")

def read_folder_contents(path):
    contents = os.listdir(path)
    return [f for f in contents if not f.startswith(".")]

class PathParser:
    def __init__(self, path):
        self.path = path
    
    def check_path_type(self):
        if os.path.isfile(self.path):
            return PathType.File
        return PathType.Directory

    def human_readable(self, delimiter):
        return os.path.basename(self.path).replace(delimiter, " ")

    def add_to_path(self, *sub):
        items = [sub_path for sub_path in sub]
        items.insert(0, self.path)
        return "/".join(items)

class GenerationParser:
    def __init__(self, line):
        self.line = line
    
    def parse(self):
        vals = self.line.split(" ")
        count = int(vals[0])
        phase = vals[1]
        fr = float(vals[2])
        igd = float(vals[3])
        hv = float(vals[4])
        arcigd = float(vals[5])
        archv = float(vals[6])

        return Generation(count, phase, fr, igd, hv, arcigd, archv)

class RunParser(PathParser):
    def __init__(self, path):
        PathParser.__init__(self, path)
    
    def parse(self):
        run_count = int(self.human_readable("-").split(" ")[-1])
        run = Run(run_count)
        with open(self.path, "r") as run_file:
            for line in run_file:
                generation_parser = GenerationParser(line)
                run.add(generation_parser.parse())
        return run

class ProblemParser(PathParser):
    def __init__(self, path):
        PathParser.__init__(self, path)
    
    def parse(self):
        name = self.human_readable("-")
        problem = Problem(name)

        for run_file in read_folder_contents(self.path):
            run_parser = RunParser(self.add_to_path(run_file))
            problem.add(run_parser.parse())

        return problem

class TestParser(PathParser):
    def __init__(self, path):
        PathParser.__init__(self, path)
    
    def parse(self):
        name = self.human_readable("-")
        test = Test(name)
        for test_suite_folder in read_folder_contents(self.path):
            path = self.add_to_path(test_suite_folder)
            for problem_folder in read_folder_contents(path):
                problem_parser = ProblemParser(self.add_to_path(test_suite_folder, problem_folder))
                test.add(problem_parser.parse())
        
        return test