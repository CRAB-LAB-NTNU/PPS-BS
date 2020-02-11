from mpl_toolkits.mplot3d import Axes3D  # noqa: F401 unused import

import matplotlib.pyplot as plt
import numpy as np
import sys, getopt


class Plotter: 

    def __init__(self, input_file, output_file, dims):
        self.input_file = input_file
        self.output_file = output_file
        self.dims = dims

        self.read_points_from_file()

    def read_points_from_file(self):
        with open(self.input_file) as f:
            lines = [[float(val) for val in line.strip().split(",")] for line in f.readlines()]
        self.xs = [coord[0] for coord in lines]
        self.ys = [coord[1] for coord in lines]
        if self.dims == 3:
            self.zs = [coord[2] for coord in lines]


    def plot(self):
        if self.dims == 2:
            self._plot2d()
        else:
            self._plot3d()

    def _plot2d(self):
        print("PLot2")
        fig = plt.figure()
        print(fig)
        ax = fig.add_axes([0,0,1,1])
        print(ax)
        ax.scatter(self.xs, self.ys, color ="r")
        ax.set_xlabel("f1")
        ax.set_ylabel("f2")
        plt.show()

    
    def _plot3d(self):
        print("Hello :)")
        fig = plt.figure()
        ax = fig.add_subplot(111,projection="3d")
        ax.scatter(self.xs,self.ys, self.zs, color="r")

        ax.set_xlabel("f1")
        ax.set_ylabel("f2")
        ax.set_zlabel("f3")
        plt.show()

def main(argv):


    input_file = ""
    output_file = ""
    dims = 0
    try:
        opts, _ = getopt.getopt(argv,"h:i:o:t:")
    except getopt.GetoptError:
        print("plotter.py -i <inputfile> -o <outputfile> -t <plot dimensions>")
        sys.exit(2)
    for opt, arg in opts:
        if opt == "-h":
            print("plotter.py -i <inputfile> -o <outputfile> -t <plot dimensions>")
            sys.exit()
        elif opt == "-i":
            input_file = arg
        elif opt == "-o":
            output_file = arg
        elif opt == "-t":
            dims = int(arg)
    
    if input_file == "":
        input_file = "defaultinput.txt"
        print("Using default inputfile:", input_file )
    if output_file == "":
        output_file = "defaultoutput.png"
        print("Using default outputfule:", output_file)
    if dims == 0:
        dims = 2
        print("Using default plot type:", dims)
        
    plotter = Plotter(input_file, output_file, dims)
    plotter.plot()


if __name__ == "__main__":
    main(sys.argv[1:])