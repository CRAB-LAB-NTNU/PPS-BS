from Parser import TestParser
from enum import Enum
import matplotlib.pyplot as plt
import os
import numpy as np
import math
import csv

class Metric(Enum):
    IGD = "igd"
    HV = "hv"
    FR = "fr"
    ARCHIVE_IGD = "archive_igd"
    ARCHIVE_HV = "archive_hv"

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
        

    def plot_values(self, metric):
        print("Creating plots for")
        [print(t.name) for t in self.tests]
        directory = self.base_directory + metric.name + "/"
        create_dir_if_missing(directory)
        for index, problem in enumerate(self.tests[0]):
            if metric is Metric.IGD:
                values = [ test[index].inverted_generational_distance() for test in self.tests ]
            elif metric is Metric.HV:
                values = [ test[index].hyper_volume() for test in self.tests ]
            elif metric is Metric.FR:
                values = [ test[index].feasibility_ratio() for test in self.tests ]
            elif metric is Metric.ARCHIVE_IGD:
                values = [ test[index].archive_inverted_generational_distance() for test in self.tests ]
            elif metric is Metric.ARCHIVE_HV:
                values = [ test[index].archive_hyper_volume() for test in self.tests ]
            
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
        self.plot_values(Metric.ARCHIVE_IGD)
        self.plot_values(Metric.ARCHIVE_HV)
        self.plot_phases()

    def mean_and_std(self):
        for index, problem in enumerate(self.tests[0]):

            with open(self.base_directory + "archive_results.csv", 'a', newline='') as csvfile:
                csvwriter = csv.writer(csvfile, delimiter=',',quotechar='|', quoting=csv.QUOTE_MINIMAL)

                median_igd = [test[index].median_archive_inverted_generational_distance(-1) for test in self.tests]
                mean_igd = [test[index].mean_archive_inverted_generational_distance(-1) for test in self.tests]
                std_igd = [test[index].std_archive_inverted_generational_distance(-1) for test in self.tests]
                median_hv = [test[index].median_archive_hyper_volume(-1) for test in self.tests]
                mean_hv = [test[index].mean_archive_hyper_volume(-1) for test in self.tests]
                std_hv = [test[index].std_archive_hyper_volume(-1) for test in self.tests]

                igd_median_row = [problem.name, "IGD", "MEDIAN"] + median_igd
                csvwriter.writerow(igd_median_row)

                igd_mean_row = ["", "", "MEAN"] + mean_igd
                csvwriter.writerow(igd_mean_row)
                
                igd_std_row = ["", "", "STD"] + std_igd
                csvwriter.writerow(igd_std_row)
                
                hv_median_row = ["", "HV", "MEDIAN"] + median_hv
                csvwriter.writerow(hv_median_row)

                hv_mean_row = ["", "", "MEAN"] + mean_hv
                csvwriter.writerow(hv_mean_row)

                hv_std_row = ["", "", "STD"] + std_hv
                csvwriter.writerow(hv_std_row)