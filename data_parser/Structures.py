import statistics

class InnerList:
    def __init__(self):
        self.container = []
    
    def add(self, item):
        self.container.append(item)

    def __getitem__(self, key):
        return self.container[key]

    def __iter__(self):
        self.n = 0
        return self

    def __next__(self):
        if self.n < len(self.container):
            r = self.container[self.n]
            self.n += 1
            return r
        else:
            raise StopIteration
    
    def size(self):
        return len(self.container)

class Generation:
    def __init__(self, count, phase, feasibility_ratio, inverted_generational_distance, hyper_volume, archive_inverted_generational_distance, archive_hyper_volume):
        self.count = count
        self.phase = phase
        self.feasibility_ratio = feasibility_ratio
        self.inverted_generational_distance = inverted_generational_distance
        self.hyper_volume = hyper_volume
        self.archive_inverted_generational_distance = archive_inverted_generational_distance
        self.archive_hyper_volume = archive_hyper_volume
    
class Run(InnerList):
    def __init__(self, run):
        InnerList.__init__(self)
        self.run = run
    
    def feasibility_ratio(self):
        return [generation.feasibility_ratio for generation in self]

    def inverted_generational_distance(self):
        return [generation.inverted_generational_distance for generation in self]
    
    def hyper_volume(self):
        return [generation.hyper_volume for generation in self]

    def archive_hyper_volume(self):
        return [generation.archive_hyper_volume for generation in self]

    def archive_inverted_generational_distance(self):
        return [generation.archive_inverted_generational_distance for generation in self]

    def binary_count(self):
        return self.phase_count("Binary")

    def push_count(self):
        return self.phase_count("Push")

    def pull_count(self):
        return self.phase_count("Pull")

    def phase_count(self, phase):
        count = 0
        for generation in self:
            if generation.phase == phase:
                count += 1
        return count

    def __str__(self):
        return str(self.run)

class Problem(InnerList):
    def __init__(self, name):
        InnerList.__init__(self)
        self.name = name
        self.archive_result = None
        
    # MEAN
    def mean_inverted_generational_distance(self, generation):
        return statistics.mean([run[generation].inverted_generational_distance for run in self])

    def mean_hyper_volume(self, generation):
        return statistics.mean([run[generation].hyper_volume for run in self])

    def mean_feasibility_ratio(self, generation):
        return statistics.mean([run[generation].feasibility_ratio for run in self])

    def mean_archive_inverted_generational_distance(self, generation):
        return statistics.mean([run[generation].archive_inverted_generational_distance for run in self])

    def mean_archive_hyper_volume(self, generation):
        return statistics.mean([run[generation].archive_hyper_volume for run in self])

    #MEDIAN
    def median_inverted_generational_distance(self, generation):
        return statistics.median([run[generation].inverted_generational_distance for run in self])
    
    def median_hyper_volume(self, generation):
        return statistics.median([run[generation].hyper_volume for run in self])

    def median_feasibility_ratio(self, generation):
        return statistics.median([run[generation].feasibility_ratio for run in self])
    
    def median_archive_inverted_generational_distance(self, generation):
        return statistics.median([run[generation].archive_inverted_generational_distance for run in self])

    def median_archive_hyper_volume(self, generation):
        return statistics.median([run[generation].archive_hyper_volume for run in self])
    
    #STD
    def std_inverted_generational_distance(self, generation):
        return statistics.stdev([run[generation].inverted_generational_distance for run in self])
    
    def std_hyper_volume(self, generation):
        return statistics.stdev([run[generation].hyper_volume for run in self])

    def std_feasibility_ratio(self, generation):
        return statistics.stdev([run[generation].feasibility_ratio for run in self])

    def std_archive_inverted_generational_distance(self, generation):
        return statistics.stdev([run[generation].archive_inverted_generational_distance for run in self])

    def std_archive_hyper_volume(self, generation):
        return statistics.stdev([run[generation].archive_hyper_volume for run in self])



    def hyper_volume(self):
        return [self.median_hyper_volume(i) for i in range(self[0].size())]
    
    def inverted_generational_distance(self):
        return [self.median_inverted_generational_distance(i) for i in range(self[0].size())]

    def feasibility_ratio(self):
        return [self.median_feasibility_ratio(i) for i in range(self[0].size())]

    def archive_inverted_generational_distance(self):
        return [self.median_archive_inverted_generational_distance(i) for i in range(self[0].size())]

    def archive_hyper_volume(self):
        return [self.median_archive_hyper_volume(i) for i in range(self[0].size())]

    def binary_count(self):
        return sum([run.binary_count() for run in self]) / self.size()

    def pull_count(self):
        return sum([run.pull_count() for run in self]) / self.size()

    def push_count(self):
        return sum([run.push_count() for run in self]) / self.size()

    def __str__(self):
        return self.name


class Test(InnerList):
    def __init__(self, name):
        InnerList.__init__(self)
        self.name = name
    
    def __str__(self):
        return self.name
