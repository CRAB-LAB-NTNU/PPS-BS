from Parser import TestParser
import matplotlib.pyplot as plt
import os

def create_dir_if_missing(path):
    if os.path.isdir(path):
        return
    print("Creating", path)
    os.makedirs(path, exist_ok=True)

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