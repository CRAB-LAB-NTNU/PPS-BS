from Parser import TestParser
from enum import Enum
import matplotlib.pyplot as plt
import os
import numpy as np
import math
class Metric(Enum):
    IGD = "igd"
    HV = "hv"
    FR = "fr"

def create_dir_if_missing(path):
    if os.path.isdir(path):
        return
    print("Creating", path)
    os.makedirs(path, exist_ok=True)

class BasePlotter:
    def __init__(self):
        self.base_directory = "graphics/graphs/"
        plt.style.use("grayscale")


class Multiplotter(BasePlotter):
    def __init__(self, paths):
        BasePlotter.__init__(self)
        self.tests = [ TestParser(path).parse() for path in paths ]
        self.base_directory += ("+".join([t.name for t in self.tests]) + "/")
        print("Creating plots for")
        [print(t.name) for t in self.tests]

    def plot_values(self, metric):
        directory = self.base_directory + metric.name + "/"
        create_dir_if_missing(directory)
        for index, problem in enumerate(self.tests[0]):
            if metric is Metric.IGD:
                values = [ test[index].inverted_generational_distance() for test in self.tests ]
            elif metric is Metric.HV:
                values = [ test[index].hyper_volume() for test in self.tests ]
            elif metric is Metric.FR:
                values = [ test[index].feasibility_ratio() for test in self.tests ]
            
            linestyles = ['-', '--', '-.', ':']
            for i, value in enumerate(values):
                plt.plot(value, label=self.tests[i].name, linestyle=linestyles[i])
            
            plt.legend(loc="upper center", bbox_to_anchor=(0.5, 1.15))
            plt.xlabel("generation")
            plt.ylabel(metric.name.upper())
            plt.savefig(directory + problem.name + ".png")
            plt.close()

    def plot_phases(self):
        base_directory = self.base_directory + "phases/"
        for test in self.tests:
            directory = base_directory + test.name + "/"
            create_dir_if_missing(directory)
            for problem in test:
                push = problem.push_count()
                pull = problem.pull_count()
                binary = problem.binary_count()
               
                phases = ["Push", "Binary", "Pull"]
                counts = [push, binary, pull]
                
                plt.barh(phases, counts)

                for index, value in enumerate(counts):
                    plt.text(value, index, str(math.floor(value)))

                plt.xlabel("generations")
                plt.savefig(directory + problem.name + ".png")
                plt.close()

    def plot(self):
        self.plot_values(Metric.IGD)
        self.plot_values(Metric.HV)
        self.plot_values(Metric.FR)
        self.plot_phases()

class Plotter:
    def __init__(self, path):
        self.test = TestParser(path).parse()
        print(self.test, "Parsed!")
        self.base_directory = "graphics/graphs/" + self.test.name + "/"
        
    def plot_mean_igd(self):
        directory = self.base_directory + "mean/igd/"
        create_dir_if_missing(directory)
        plt.xlabel = "generation"
        plt.ylabel = "IGD"
        for problem in self.test:
            self.plot_and_save(problem.inverted_generational_distance(), directory + problem.name + ".png")
    
    def plot_mean_hv(self):
        directory = self.base_directory + "mean/hv/"
        create_dir_if_missing(directory)
        plt.xlabel = "generation"
        plt.ylabel = "HV"
        for problem in self.test:
            self.plot_and_save(problem.hyper_volume(), directory + problem.name + ".png")

    def plot_mean_fr(self):
        directory = self.base_directory + "mean/fr/"
        create_dir_if_missing(directory)
        plt.xlabel = "generation"
        plt.ylabel = "FR"
        for problem in self.test:
            self.plot_and_save(problem.feasibility_ratio(), directory + problem.name + ".png")

    def plot_all_plot_runs_hv(self):
        root_directory = self.base_directory + "runs/hv/"
        plt.xlabel = "generations"
        plt.ylabel = "HV"
        for problem in self.test:
            directory = root_directory + problem.name + "/"
            create_dir_if_missing(directory)
            for run in problem:
                self.plot_and_save(run.hyper_volume(), directory + str(run.run) + ".png")

    def plot_all_plot_runs_fr(self):
        root_directory = self.base_directory + "runs/fr/"
        plt.xlabel = "generations"
        plt.ylabel = "FR"
        for problem in self.test:
            directory = root_directory + problem.name + "/"
            create_dir_if_missing(directory)
            for run in problem:
                self.plot_and_save(run.feasibility_ratio(), directory + str(run.run) + ".png")

    def plot_all_plot_runs_igd(self):
        root_directory = self.base_directory + "runs/igd/"
        plt.xlabel = "generations"
        plt.ylabel = "IGD"
        for problem in self.test:
            directory = root_directory + problem.name + "/"
            create_dir_if_missing(directory)
            for run in problem:
                self.plot_and_save(run.inverted_generational_distance(), directory + str(run.run) + ".png")

    def plot_and_save(self, values, path):
        plt.plot(values)
        plt.savefig(path)
        plt.close()

    def plot_mean(self):
        self.plot_mean_fr()
        self.plot_mean_hv()
        self.plot_mean_igd()

    def plot(self):
        self.plot_all_plot_runs_fr()
        self.plot_all_plot_runs_hv()
        self.plot_all_plot_runs_igd()
        self.plot_mean()