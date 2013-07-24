__author__ = 'cody'
from bulbs.model import Node, Relationship
from bulbs.property import String, Integer, DateTime
from bulbs.utils import current_datetime
from bulbs.neo4jserver import Graph
from bulbs.neo4jserver import Graph, Config, NEO4J_URI


config = Config(NEO4J_URI, "james", "secret")
g = Graph()

tomCruise = g.vertices.create(name="Tom")
missionImpossible = g.vertices.create(name="Mission")
g.edges.create(tomCruise, "knows", missionImpossible)

dict_object = {'Actor': "Tom Cruise", 'Movie': "Mission Impossible", 'Director': "Brian De Palma",
               'Actors': ["Jon Voight", "Emmanuelle Beart", "Henry Czerny", "Jean Reno", "Ving Rhames"]}
for key, value in dict_object.items():
    if type(value) is list:
        for item in value:
            print(key, item)
    else:
        print(key, value)

